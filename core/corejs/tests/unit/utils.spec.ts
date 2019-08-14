import {
	newFormWithStates,
	setPushState,
	setFormValue,
} from '@/utils';


describe('utils', () => {
	it('newFormWithStates', () => {
		const fd = newFormWithStates({ f1: ['1'], f2: ['2'] });
		expect(fd.get('f1')).toBe('1');
	});

	it('setFormValue, getFormValue, getFormValueAsArray', () => {
		const fd = new FormData();
		setFormValue(fd, 'f1', ['1', '2']);
		expect(fd.getAll('f1')).toEqual(['1', '2']);
		setFormValue(fd, 'f1', '1');
		expect(fd.getAll('f1')).toEqual(['1']);
		expect(fd.get('f1')).toEqual('1');
	});

	it('setPushState', () => {

		const pusher = {
			pushed: {} as any,
			pushState: (data: any, title: string, url?: string | null) => {
				pusher.pushed = {
					data,
					title,
					url,
				};
			},
		};


		const { newEventFuncId, eventURL } = setPushState(
			{
				id: 'hello',
				pushState: { name: 'felix' },
			},
			'/page1?hello=1&page=2',
			pusher,
			false,
		);
		expect(eventURL).toBe('/page1?__execute_event__=hello&hello=1&name=felix&page=2');
		expect(pusher.pushed.url).toBe('/page1?hello=1&name=felix&page=2');
		expect(newEventFuncId.pushState).toEqual({ name: ['felix'] });

		const r2 = setPushState(
			{
				id: 'hello',
				pushState: 'name=felix',
			},
			'/page1?hello=1&page=2',
			pusher,
			false,
		);
		expect(r2.eventURL).toBe('/page1?__execute_event__=hello&name=felix');
		expect(pusher.pushed.url).toBe('/page1?name=felix');
		expect(pusher.pushed.data).toEqual({ name: 'felix', url: '/page1?name=felix' });
		expect(r2.newEventFuncId.pushState).toEqual({ name: ['felix'] });

	});
});
