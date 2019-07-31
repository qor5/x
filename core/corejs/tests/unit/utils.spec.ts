import {
	newFormWithStates,
	setPushState,
	PushStateFunc,
	setFormValue,
	getFormValue,
	getFormValueAsArray,
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
		expect(getFormValueAsArray(fd, 'f1')).toEqual(['1', '2']);
		setFormValue(fd, 'f1', '1');
		expect(fd.getAll('f1')).toEqual(['1']);
		expect(fd.get('f1')).toEqual('1');
		expect(getFormValue(fd, 'f1')).toEqual('1');
		expect(getFormValueAsArray(fd, 'f1')).toEqual(['1']);
	});

	it('setPushState', () => {

		let pushed: any;
		const pushFunc: PushStateFunc = (data: any, title: string, url?: string | null) => {
			pushed = {
				data,
				title,
				url,
			};
		};

		const search = setPushState(
			{ name: 'felix' },
			'hello=1&page=2',
			'/page1',
			pushFunc,
		);
		expect(search).toBe('&hello=1&name=felix&page=2');
		expect(pushed.url).toBe('/page1?hello=1&name=felix&page=2');
	});
});
