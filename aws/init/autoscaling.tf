resource "aws_autoscaling_group" "_" {
  count = 0
  name                 = "${var.ns}"
  desired_capacity     = 2
  health_check_type    = "EC2"
  launch_configuration = "${aws_launch_configuration._.name}"
  max_size             = 4
  min_size             = 2
  force_delete         = true

  vpc_zone_identifier  = [
    "${aws_subnet.priv_d.id}",
    "${aws_subnet.priv_c.id}",
  ]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_launch_configuration" "_" {
  count = 0
  name_prefix                 = "${var.ns}-"
  associate_public_ip_address = true
  iam_instance_profile        = "${aws_iam_instance_profile._.id}"
  image_id                    = "${data.aws_ami.ecs_optimized.id}"
  instance_type               = "t2.micro"
  security_groups             = ["${aws_security_group.ecs.id}", "${aws_security_group.svc.id}"]
  user_data                   = "#!/bin/bash\necho ECS_CLUSTER=${aws_ecs_cluster._.name} >> /etc/ecs/ecs.config"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = 8
    delete_on_termination = true
  }

  lifecycle {
    create_before_destroy = true
  }
}

data "aws_ami" "ecs_optimized" {
  most_recent = true
  owners = ["amazon"] # AWS

  filter {
    name   = "name"
    values = ["amzn-ami*amazon-ecs-optimized"]
  }
  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

resource "aws_iam_instance_profile" "_" {
  count = 0
  name  = "${var.ns}"
  role  = "${aws_iam_role.ec2.name}"
}
