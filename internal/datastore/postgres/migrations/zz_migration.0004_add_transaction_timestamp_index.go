package migrations

import "context"

const createIndexOnTupleTransactionTimestamp = `
	CREATE INDEX ix_relation_tuple_transaction_by_timestamp on relation_tuple_transaction(timestamp);
`

func init() {
	if err := DatabaseMigrations.Register("add-transaction-timestamp-index", "add-unique-living-ns", func(ctx context.Context, apd *AlembicPostgresDriver) error {
		_, err := apd.db.Exec(ctx, createIndexOnTupleTransactionTimestamp)
		return err
	}); err != nil {
		panic("failed to register migration: " + err.Error())
	}
}
