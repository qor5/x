import {encodeFilterData} from '@/components/FilterData';
import moment from 'moment';
import qs from 'query-string';

describe('filter', () => {
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
					timezone: 'utc',
				},
				{
					key: 'updated',
					label: 'Updated',
					itemType: 'DateItem',
					selected: true,
					modifier: 'equals',
					valueIs: '2018-04-09',
				},
			])).toEqual('created=1523232000&updated=1523203200');
		});

		it('inTheLast days', () => {
			const v = encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'inTheLast',
					inTheLastValue: '3',
					inTheLastUnit: 'days',
				},
			]);

			expect(qs.parse(v)).toMatchObject({
				'created.lt': expect.stringMatching(/.+/),
				'created.gte': expect.stringMatching(/.+/),
			});
		});

		it('inTheLast months', () => {
			const v = encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'inTheLast',
					inTheLastValue: '3',
					inTheLastUnit: 'months',
				},
			]);

			expect(qs.parse(v)).toMatchObject({
				'created.lt': expect.stringMatching(/.+/),
				'created.gte': expect.stringMatching(/.+/),
			});
		});

		it('isAfterOrOn', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'isAfterOrOn',
					valueIs: moment('2018-04-09'),
					timezone: 'utc',
				},
			])).toEqual('created.gte=1523232000');
		});

		it('isAfter', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'isAfter',
					valueIs: moment('2018-04-09'),
				},
			])).toEqual('created.gt=1523289599');
		});

		it('isBefore', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'isBefore',
					valueIs: moment('2018-04-09'),
					timezone: 'utc',
				},
			])).toEqual('created.lt=1523232000');
		});

		it('isBeforeOrOn', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'isBeforeOrOn',
					valueIs: moment('2018-04-09'),
					timezone: 'utc',
				},
			])).toEqual('created.lte=1523318399');
		});

		it('between', () => {
			expect(encodeFilterData([
				{
					key: 'created',
					label: 'Created',
					itemType: 'DateItem',
					selected: true,
					modifier: 'between',
					valueFrom: moment('2018-04-09'),
					valueTo: moment('2018-04-10'),
				},
				{
					key: 'created1',
					label: 'Created1',
					itemType: 'DateItem',
					selected: true,
					modifier: 'between',
					valueFrom: moment('2018-04-09'),
				},
				{
					key: 'created2',
					label: 'Created2',
					itemType: 'DateItem',
					selected: true,
					modifier: 'between',
					valueTo: moment('2018-04-09'),
				},
				{
					key: 'created3',
					label: 'Created3',
					itemType: 'DateItem',
					selected: true,
					modifier: 'between',
				},
				{
					key: 'confirmed',
					label: 'Confirmed',
					itemType: 'DateItem',
					selected: true,
					modifier: 'between',
					valueFrom: moment('2018-04-09'),
					valueTo: moment('2018-04-10'),
					timezone: 'utc',
				},
			])).toEqual(
				// tslint:disable-next-line: max-line-length
				'created.gte=1523203200&created.lt=1523376000&created1.gte=1523203200&created2.lt=1523289600&confirmed.gte=1523232000&confirmed.lt=1523404800',
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
					valuesAre: ['2','3','7'],
				},
				{
					key: 'group_channel',
					label: 'Group&Channel',
					itemType: 'LinkageSelectItem',
					selected: true,
					modifier: 'equals',
					valuesAre: ['1','2'],
				},
			])).toEqual('province_city_district=2%2C3%2C7&group_channel=1%2C2');
		});
	});

});
