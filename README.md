# expense-tracker
A CLI tool for personal expense tracking.

## Add Expense
Command to add expenses.
```shell
./extr add -d "description of the expense" -a <amount> -c <categories>
```
Following are the list of supported categories:
- groceries
- transport
- medical
- rent
- entertainment
- uncategorised

If no valid category is provided, it is set to `uncategorised`.

## List Expenses
Command to list existing expense entries.
- This lists all the expense entries
```shell
./extr list
```
- To list expenses of a month this year.
```shell
./extr list -m <month_number>
```
- To list expenses of a particular month-year.
```shell
./extr list -m <month_number> -y <year>
```
- To list categorical expenses.
```shell
./extr list -c <comma,sperated,categories>
```
Note: Only yearly expenses are listed. If year is not passed in the flags, expenses are listed for the current year.

## Update Expense
Command to udpate existing expense entry.
```shell
./extr update --id <expense_id> -d "description of the expense" -a <amount> -c <categories>
```

## Delete Expense
Command to delete existing expense entry.
```shell
./extr delete --id <expense_id>
```

## Summary
By default year is taken to be the current one.
- Fetch total expenses of the current month.
```shell
./extr summary
```
- Fetch total expenses of an year.
```shell
./extr summary -y <year>
```
- Fetch total expenses of another month.
```shell
./extr summary -m <month>
```
- Fetch total expenses of month another year.
```shell
./extr summary -m <month> -y <year>
```

# Budget
Bugetting is also supported for the expense tracker.
A user can set a monthly budget.
If the total expense of the month goes beyond the budget.
User gets warning after each new expense being added.

## Setting a budget
- If no month or year is mentioned, then the budget is set for the current month.
```shell
./extr add budget -a <amount>
```
- To set budget for any other month in current or some other year user can pass then as flags.
```shell
./extr add budget -a <amount> -m <month_number> -y <year>
```

## List budget
- To view budget for the current month, user can just run.
```shell
./extr add budget
```
- To list all the budget entries.
```shell
./extr add budget list
```
- To list particular budget entries, user can provide the month and the year.
```shell
./extr add budget list -m <month_number> -y <year>
```

## Update budget
- To update budget entry for the current month.
```shell
./extr update budget update -a <amount>
```
- To update budget entry for a particular month. If year is not passed it is assumed to be current one.
```shell
./extr update budget update -a <amount> -m <month> -y <year>
```

## Delete budget
- To delete budget for the current month.
```shell
./extr update budget delete
```
- To delete budget for a particular month. If year is not passed it is assumed to be the current one.
```shell
./extr update budget delete -m <month_number> -y <year>
```

## Export Expense data
- A user can export the expense data to a CSV file using the following command.
If no file path is passed the exported data is saved as `expense.csv` in the current directory.
```shell
./extr export -o <csv_file_path>
```
