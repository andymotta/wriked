FROM scratch
EXPOSE 8080
ENTRYPOINT ["/wriked"]
COPY ./bin/ /