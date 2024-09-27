
export type Supply = {
	name: string;
	description: any;
	photo: any;
	branch_supplies_stock: any;
	inventory_supplies_stock: any;
};
export type Article = {
	name: string;
	description: any;
	photo: any;
	branch_articles_stock: any;
	inventory_articles_stock: any;
};
export type Safebox = {
	cents10: any;
	cents50: any;
	coins1: any;
	coins2: any;
	coins5: any;
	coins10: any;
	coins20: any;
	bills20: any;
	bills50: any;
	bills100: any;
	bills200: any;
	bills500: any;
	bills1000: any;
	branch_safeboxes: any;
	turn_safeboxes: any;
	inventory_safeboxes: any;
	action_safebox: any;
};
export type Income = {
	reason: string;
	income: number;
	sale_incomes: any;
};
export type Expense = {
	reason: string;
	expense: number;
	sale_expenses: any;
};
export type Argument = {
	complaint: boolean;
	score: any;
	argument: string;
	sale_arguments: any;
};