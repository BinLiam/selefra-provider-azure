package sql

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlServerAdminsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlServerAdminsGenerator{}

func (x *TableAzureSqlServerAdminsGenerator) GetTableName() string {
	return "azure_sql_server_admins"
}

func (x *TableAzureSqlServerAdminsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlServerAdminsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlServerAdminsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureSqlServerAdminsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(*armsql.Server)
			cl := client.(*azure_client.Client)
			svc, err := armsql.NewServerAzureADAdministratorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			group, err := azure_client.ParseResourceGroup(*p.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListByServerPager(group, *p.Name, nil)
			for pager.More() {
				p, err := pager.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- p.Value
			}
			return nil
		},
	}
}

func (x *TableAzureSqlServerAdminsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureSqlServerAdminsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_servers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_servers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
	}
}

func (x *TableAzureSqlServerAdminsGenerator) GetSubTables() []*schema.Table {
	return nil
}
