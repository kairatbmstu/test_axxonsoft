#!/bin/bash

# Set PostgreSQL configuration
export PG_HOST=localhost
export PG_PORT=5432
export PG_USERNAME=test_axxonsoft
export PG_PASSWORD=123456
export PG_DATABASE=test_axxonsoft

# Set RabbitMQ configuration
export RABBITMQ_HOST=localhost
export RABBITMQ_PORT=15672
export RABBITMQ_USERNAME=guest
export RABBITMQ_PASSWORD=guest

# Optional: Display the environment variables
echo "PostgreSQL Configuration:"
echo "PG_HOST: $PG_HOST"
echo "PG_PORT: $PG_PORT"
echo "PG_USERNAME: $PG_USERNAME"
echo "PG_PASSWORD: $PG_PASSWORD"
echo "PG_DATABASE: $PG_DATABASE"

echo "RabbitMQ Configuration:"
echo "RABBITMQ_HOST: $RABBITMQ_HOST"
echo "RABBITMQ_PORT: $RABBITMQ_PORT"
echo "RABBITMQ_USERNAME: $RABBITMQ_USERNAME"
echo "RABBITMQ_PASSWORD: $RABBITMQ_PASSWORD"