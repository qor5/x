import {encodeFilterData} from '@/components/FilterData';

describe('filter', () => {
	describe('encodeFilterData DatetimeRangeItem', () => {
		it('between', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DatetimeRangeItem',
					selected: true,
					modifier: 'between',
					valueFrom: '2018-04-09 00:00',
					valueTo: '2018-04-10 00:00',
				},
				{
					key: 'created1',
					label: 'Created1',
					itemType: 'DatetimeRangeItem',
					selected: true,
					modifier: 'between',
					valueFrom: '2018-04-09 00:00',
				},
				{
					key: 'created2',
					label: 'Created2',
					itemType: 'DatetimeRangeItem',
					selected: true,
					modifier: 'between',
					valueTo: '2018-04-09 00:00',
				},
				{
					key: 'created3',
					label: 'Created3',
					itemType: 'DatetimeRangeItem',
					selected: true,
					modifier: 'between',
				},
				{
					key: 'confirmed',
					label: 'Confirmed',
					itemType: 'DatetimeRangeItem',
					selected: true,
					modifier: 'between',
					valueFrom: '2018-04-09 00:00',
					valueTo: '2018-04-10 00:00',
				},
			])).toEqual(
				// tslint:disable-next-line: max-line-length
				'created.gte=2018-04-09%2000%3A00&created.lt=2018-04-10%2000%3A00&created1.gte=2018-04-09%2000%3A00&created2.lt=2018-04-09%2000%3A00&confirmed.gte=2018-04-09%2000%3A00&confirmed.lt=2018-04-10%2000%3A00',
			);
		});
	});

	describe('encodeFilterData DateRangeItem', () => {
		it('equals', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateRangeItem',
					selected: true,
					modifier: 'between',
					valueFrom: '2019-09-10',
					valueTo: '2019-09-20',
				},
			])).toEqual(
				'created.gte=2019-09-10&created.lte=2019-09-20',
			);
		});
	});

	describe('encodeFilterData DateItem', () => {
		it('equals', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'equals',
					valueIs: '2018-04-09',
				},
			])).toEqual(
				'created=2018-04-09',
			);
		});
	});

	describe('encodeFilterData NumberItem', () => {
		it('equals', () => {
			expect(encodeFilterData([
				{
					key: 'amount',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'equals',
					valueIs: '12',
				},
				{
					key: 'amount1',
					label: 'Amount1',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'equals',
					valueIs: null,
				},
			])).toEqual('amount=12');
		});

		it('between', () => {
			expect(encodeFilterData([
				{
					key: 'amount',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'between',
					valueFrom: 12,
					valueTo: 24,
				},
				{
					key: 'amount1',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'between',
					valueTo: 24,
				},
				{
					key: 'amount2',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'between',
					valueFrom: 12,
				},
				{
					key: 'amount3',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'between',
				},
			])).toEqual('amount.gte=12&amount.lte=24&amount1.lte=24&amount2.gte=12');
		});

		it('greaterThan', () => {
			expect(encodeFilterData([
				{
					key: 'amount',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'greaterThan',
					valueIs: 12,
				},
			])).toEqual('amount.gt=12');
		});

		it('lessThan', () => {
			expect(encodeFilterData([
				{
					key: 'amount',
					label: 'Amount',
					itemType: 'NumberItem',
					selected: true,
					modifier: 'lessThan',
					valueIs: 12,
				},
			])).toEqual('amount.lt=12');
		});
	});


	describe('encodeFilterData StringItem', () => {
		it('equals', () => {
			expect(encodeFilterData([
				{
					key: 'name',
					label: 'Name',
					itemType: 'StringItem',
					selected: true,
					modifier: 'equals',
					valueIs: 'felix[]',
				},
				{
					key: 'name1',
					label: 'Name',
					itemType: 'StringItem',
					selected: true,
					modifier: 'equals',
					valueIs: null,
				},
			])).toEqual('name=felix%5B%5D');
		});

		it('contains', () => {
			expect(encodeFilterData([
				{
					key: 'name',
					label: 'Name',
					itemType: 'StringItem',
					selected: true,
					modifier: 'contains',
					valueIs: 'felix[]',
				},
				{
					key: 'name',
					label: 'Name',
					itemType: 'StringItem',
					selected: true,
					modifier: 'contains',
				},
			])).toEqual('name.ilike=felix%5B%5D');
		});
	});

	describe('encodeFilterData LinkageSelectItem', () => {
		it('equals', () => {
			expect(encodeFilterData([
				{
					key: 'province_city_district',
					label: 'Province&City&District',
					itemType: 'LinkageSelectItem',
					selected: true,
					modifier: 'equals',
					valuesAre: ['2', '3', '7'],
				},
				{
					key: 'group_channel',
					label: 'Group&Channel',
					itemType: 'LinkageSelectItem',
					selected: true,
					modifier: 'equals',
					valuesAre: ['1', '2'],
				},
			])).toEqual('province_city_district=2%2C3%2C7&group_channel=1%2C2');
		});
	});

});
