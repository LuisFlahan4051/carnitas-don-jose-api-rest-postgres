
export type UserSafeboxAction = {
	withdrawal: boolean;
	user_id: any;
};
export type ActionSafebox = {
	safebox_action_id: any;
	safebox_id: any;
};
export type AdminNotification = {
	type: string;
	solved: boolean;
	description: string;
	branch_id: any;
	user_id: any;
	images: any;
};
export type AdminNotificationImage = {
	image: string;
	notification_id: any;
};
export type ServerLogs = {
	id: any;
	created_at: any;
	transaction: string;
	user_id: any;
	branch_id: any;
	root: any;
};
export type Pagination = {
	page: any;
	today: any;
	since: any;
	to: any;
};
export type LoginForm = {
	username: string;
	password: string;
};