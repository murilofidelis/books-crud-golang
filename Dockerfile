FROM gcr.io/distroless/base-debian11

COPY api ./
COPY .env ./

EXPOSE 8082

ENTRYPOINT ["/api"]
