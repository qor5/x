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


		const search = setPushState(
			{ name: 'felix' },
			'hello=1&page=2',
			'/page1',
			pusher,
		);
		expect(search).toBe('&hello=1&name=felix&page=2');
		expect(pusher.pushed.url).toBe('/page1?hello=1&name=felix&page=2');
	});
});
