export interface ValueOp {
	value: string | string[];
	add?: boolean;
	remove?: boolean;
}

interface PushState {
	[key: string]: null | undefined | string | string[] | ValueOp;
}

export interface EventFuncID {
	id: string;
	params?: string[];
	pushState?: PushState;
}

export interface PortalUpdate {
	name: string;
	schema: string;
	afterLoaded?: string;
}

export interface EventResponse {
	states?: any;
	schema?: any;
	data?: any;
	redirectURL?: string;
	pageTitle?: string;
	reloadWindowURL?: string;
	reload: boolean;
	reloadPortals?: string[];
	updatePortals?: PortalUpdate[];
}

export interface StatePusher {
	pushState(data: any, title: string, url?: string | null): void;
}
