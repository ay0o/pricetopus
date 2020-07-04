provider "aws" {
  region = "eu-west-1"
}

data "aws_vpc" "vpc" {
  default = true
}

data "aws_subnet_ids" "subnets" {
  vpc_id = data.aws_vpc.vpc.id
}

resource "aws_ecs_cluster" "cluster" {
  name               = "pricetopus"
  capacity_providers = ["FARGATE_SPOT"]
  default_capacity_provider_strategy {
    capacity_provider = "FARGATE_SPOT"
  }
}

resource "aws_ecs_task_definition" "task" {
  family                   = "pricetopus"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  container_definitions    = <<DEFINITION
    [
        {
            "environment": [
                {"name": "PRICETOPUS_EMAIL_SERVER", "value": "${var.server}"},
                {"name": "PRICETOPUS_EMAIL_SERVER_PORT", "value": "${var.port}"},
                {"name": "PRICETOPUS_EMAIL_USER", "value": "${var.sender}"},
                {"name": "PRICETOPUS_EMAIL_PASSWORD", "value": "${var.password}"},
                {"name": "PRICETOPUS_EMAIL_TO", "value": "${var.recipient}"},
                {"name": "PRICETOPUS_PRODUCT_URL", "value": ""},
                {"name": "PRICETOPUS_PRODUCT_PRICE", "value": ""}
            ],
            "essential": true,
            "image": "ay0o/pricetopus",
            "memory": 128,
            "name": "pricetopus"
        }
    ]
    DEFINITION
}

resource "aws_iam_role" "ecs_events" {
  name = "ecs_events"

  assume_role_policy = <<DOC
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "events.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
DOC
}

resource "aws_iam_role_policy" "ecs_events_run_task_with_any_role" {
  name = "ecs_events_run_task_with_any_role"
  role = aws_iam_role.ecs_events.id

  policy = <<DOC
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "iam:PassRole",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "ecs:RunTask",
            "Resource": "${aws_ecs_task_definition.task.arn}"
        }
    ]
}
DOC
}

resource "aws_cloudwatch_event_rule" "trigger" {
  name                = "pricetopus_trigger"
  schedule_expression = var.trigger
}

resource "aws_cloudwatch_event_target" "cron" {
  for_each  = var.products
  target_id = each.key
  arn       = aws_ecs_cluster.cluster.arn
  rule      = aws_cloudwatch_event_rule.trigger.name
  role_arn  = aws_iam_role.ecs_events.arn

  ecs_target {
    task_count          = 1
    task_definition_arn = aws_ecs_task_definition.task.arn
    launch_type         = "FARGATE"
    network_configuration {
      subnets          = data.aws_subnet_ids.subnets.ids
      assign_public_ip = true
    }
  }

  input = <<OVERRIDE
{
    "containerOverrides": [
        {
            "name": "pricetopus",
            "environment": [
                {
                    "name": "PRICETOPUS_PRODUCT_URL",
                    "value": "${each.value["url"]}"
                },
                {
                    "name": "PRICETOPUS_PRODUCT_PRICE",
                    "value": "${each.value["price"]}"
                }
            ]
        }
    ]
}
OVERRIDE
}