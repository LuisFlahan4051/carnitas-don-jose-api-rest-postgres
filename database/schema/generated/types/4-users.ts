
export type Role = {
	name: string;
	access_level: any;
	branch_user_roles: any;
	turn_user_roles: any;
	inherit_user_roles: any;
};
export type User = {
	name: any;
	lastname: any;
	username: string;
	password: string;
	photo: any;
	verified: any;
	warning: any;
	darktheme: any;
	active_contract: any;
	address: any;
	born: any;
	degree_study: any;
	relation_ship: any;
	curp: any;
	rfc: any;
	citizen_id: any;
	credential_id: any;
	origin_state: any;
	score: any;
	qualities: any;
	defects: any;
	origin_branch_id: any;
	branch_id: any;
	inherit_user_roles: any;
	user_phones: any;
	user_mails: any;
	user_reports: any;
	monetary_bounds: any;
	monetary_discounts: any;
	branch_user_roles: any;
	turns: any;
	turn_user_roles: any;
	sales: any;
	user_safebox_actions: any;
	admin_notifications: any;
};
export type InheritUserRole = {
	role_id: any;
	user_id: any;
};
export type UserPhone = {
	phone: string;
	main: any;
	user_id: any;
};
export type UserMail = {
	mail: string;
	main: any;
	user_id: any;
};
export type UserReport = {
	reason: string;
	user_id: any;
};
export type MonetaryBound = {
	reason: string;
	bound: number;
	user_id: any;
};
export type MonetaryDiscount = {
	reason: string;
	discount: number;
	user_id: any;
};
export type BranchUserRole = {
	branch_id: any;
	user_id: any;
	role_id: any;
};