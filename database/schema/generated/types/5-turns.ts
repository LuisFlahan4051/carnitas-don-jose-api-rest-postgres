
export type Turn = {
	start_date: any;
	end_date: any;
	active: any;
	incomes_counter: number;
	net_incomes_counter: number;
	expenses_counter: number;
	user_id: any;
	branch_id: any;
	turn_user_roles: any;
	safebox_received: any;
	safebox_finished: any;
	sales: any;
	inventories: any;
};
export type TurnUserRole = {
	login_date: any;
	logout_date: any;
	user_id: any;
	turn_id: any;
	role_id: any;
};
export type TurnSafebox = {
	turn_id: any;
	safebox_id: any;
};