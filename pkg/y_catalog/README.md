### Parquet Schema vs Iceberg Schema

In Apache Iceberg, you will use the Iceberg table schema rather than the Parquet schema. Iceberg is a table format that
provides additional features and capabilities on top of Parquet, such as schema evolution, transactional writes, and
time travel. Iceberg maintains its own metadata and schema information, allowing for easier management and evolution of
the table schema.

> Without schema evolution, you can read schema from one parquet file, and while reading rest of files assume it stays the
same.

### Reference
- [immunodb](https://github.com/codenotary/immudb/blob/65a25a5b71de2522d93ea0c5f5fc585c9a7a9f69/embedded/sql/catalog.go#L81)
