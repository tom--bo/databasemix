#!/bin/bash

# PostgreSQL 16 Test Container Management Script

CONTAINER_NAME="dbmix-postgres16-test"
COMPOSE_FILE="docker-compose.yml"

case "$1" in
    start)
        echo "Starting PostgreSQL 16 test container..."
        docker compose -f $COMPOSE_FILE up -d
        echo "Waiting for PostgreSQL 16 to be ready..."

        until docker exec $CONTAINER_NAME pg_isready -U postgres; do
            echo "Waiting for PostgreSQL 16..."
            sleep 2
        done

        echo "PostgreSQL 16 is ready!"
        echo "Connection details:"
        echo "  Host: localhost"
        echo "  Port: 5416"
        echo "  User: postgres"
        echo "  Password: rootpass"
        echo "  Database: testdb"
        echo ""
        echo "Test command:"
        echo "  ./databasemix -type=postgres -host=localhost -port=5416 -user=postgres -password=rootpass -database=testdb"
        ;;

    stop)
        echo "Stopping PostgreSQL 16 test container..."
        docker compose -f $COMPOSE_FILE down
        ;;

    restart)
        echo "Restarting PostgreSQL 16 test container..."
        docker compose -f $COMPOSE_FILE down
        docker compose -f $COMPOSE_FILE up -d
        ;;

    logs)
        docker compose -f $COMPOSE_FILE logs -f
        ;;

    shell)
        echo "Opening PostgreSQL shell..."
        docker exec -it $CONTAINER_NAME psql -U postgres testdb
        ;;

    status)
        docker compose -f $COMPOSE_FILE ps
        ;;

    clean)
        echo "Cleaning up PostgreSQL 16 test container and volumes..."
        docker compose -f $COMPOSE_FILE down -v
        docker system prune -f
        ;;

    *)
        echo "Usage: $0 {start|stop|restart|logs|shell|status|clean}"
        echo ""
        echo "Commands:"
        echo "  start   - Start the PostgreSQL 16 container"
        echo "  stop    - Stop the PostgreSQL 16 container"
        echo "  restart - Restart the PostgreSQL 16 container"
        echo "  logs    - Show container logs"
        echo "  shell   - Open PostgreSQL shell"
        echo "  status  - Show container status"
        echo "  clean   - Remove container and volumes"
        exit 1
        ;;
esac
