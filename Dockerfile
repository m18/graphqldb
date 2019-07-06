FROM postgres

ENV POSTGRES_USER store
ENV POSTGRES_PASSWORD p
ENV POSTGRES_DB store

COPY init.sql /docker-entrypoint-initdb.d
