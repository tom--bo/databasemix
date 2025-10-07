FROM mysql:8.0

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=rootpass
ENV MYSQL_DATABASE=testdb
ENV MYSQL_USER=testuser
ENV MYSQL_PASSWORD=testpass

# Expose MySQL port
EXPOSE 3380

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=30s --retries=3 \
  CMD mysqladmin ping -h localhost -u root -p$MYSQL_ROOT_PASSWORD || exit 1