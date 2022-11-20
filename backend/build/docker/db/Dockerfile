FROM mysql:latest
# RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8
# ENV LANG ru_RU.UTF8

COPY migr.sql /docker-entrypoint-migrdb.d/10-migr.sql