
FROM php:8.4-fpm-alpine

# Install necessary packages including icu-dev
RUN apk update && apk add --no-cache \
    zip unzip curl libpng-dev oniguruma-dev libxml2-dev libzip-dev \
    mariadb-client \
    icu-dev \
    autoconf g++ make \
    && docker-php-ext-install pdo_mysql zip

WORKDIR /var/www/html

COPY ./app /var/www/html

# Set permissions for storage and cache directories
RUN chown -R www-data:www-data /var/www/html
RUN chmod -R 775 /var/www/html

# Switch to the www-data user
USER www-data

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

RUN cp .env.example .env

RUN composer install --no-dev --optimize-autoloader

RUN php artisan key:generate