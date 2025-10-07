#!/bin/bash

# MySQL 8.4 Test Container Management Script

CONTAINER_NAME="dbmix-mysql84-test"
COMPOSE_FILE="docker-compose.yml"

case "$1" in
    start)
        echo "Starting MySQL 8.4 test container..."
        docker-compose -f $COMPOSE_FILE up -d
        echo "Waiting for MySQL 8.4 to be ready..."

        # Wait for MySQL to be ready
        until docker exec $CONTAINER_NAME mysqladmin ping -h localhost -u root -prootpass --silent; do
            echo "Waiting for MySQL 8.4..."
            sleep 2
        done

        echo "MySQL 8.4 is ready!"
        echo "Connection details:"
        echo "  Host: localhost"
        echo "  Port: 3384"
        echo "  User: root"
        echo "  Password: rootpass"
        echo "  Database: testdb"
        echo ""
        echo "Test command:"
        echo "  ./dbmix -host=localhost -port=3384 -user=root -password=rootpass -database=testdb"
        ;;

    stop)
        echo "Stopping MySQL 8.4 test container..."
        docker-compose -f $COMPOSE_FILE down
        ;;

    restart)
        echo "Restarting MySQL 8.4 test container..."
        docker-compose -f $COMPOSE_FILE down
        docker-compose -f $COMPOSE_FILE up -d
        ;;

    logs)
        docker-compose -f $COMPOSE_FILE logs -f
        ;;

    shell)
        echo "Opening MySQL shell..."
        docker exec -it $CONTAINER_NAME mysql -u root -prootpass testdb
        ;;

    status)
        docker-compose -f $COMPOSE_FILE ps
        ;;

    clean)
        echo "Cleaning up MySQL 8.4 test container and volumes..."
        docker-compose -f $COMPOSE_FILE down -v
        docker system prune -f
        ;;

    *)
        echo "Usage: $0 {start|stop|restart|logs|shell|status|clean}"
        echo ""
        echo "Commands:"
        echo "  start   - Start the MySQL 8.4 container"
        echo "  stop    - Stop the MySQL 8.4 container"
        echo "  restart - Restart the MySQL 8.4 container"
        echo "  logs    - Show container logs"
        echo "  shell   - Open MySQL shell"
        echo "  status  - Show container status"
        echo "  clean   - Remove container and volumes"
        exit 1
        ;;
esac