FROM cockroachdb/cockroach

COPY deploying_monolith_linksrus/linksrus/linkgraph/store/cdb/migrations /migrations
COPY Chapter10/cdb-schema/bootstrap-db.sh .

ENTRYPOINT ["bash", "./bootstrap-db.sh"]
