/*
import { localTimezoneAbbr } from "@/components/Filter";
import { encodeFilterData } from "@/components/FilterData";
import moment from "moment";
import qs from "query-string";

describe("filter", () => {
	it("localTimezoneAbbr", () => {
		expect(localTimezoneAbbr()).toEqual("CST");
	});

	describe("encodeFilterData DateItem", () => {
		it("equals", () => {
			expect(encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "equals",
					value: moment("2018-04-09"),
					timezone: "utc"
				},
				{
					key: "updated",
					label: "Updated",
					itemType: "DateItem",
					selected: true,
					modifier: "equals",
					value: moment("2018-04-09")
				}
			])).toEqual("created=1523232000&updated=1523203200");
		});

		it("inTheLast days", () => {
			var v = encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "inTheLast",
					inTheLastValue: "3",
					inTheLastUnit: "days"
				}
			]);

			expect(qs.parse(v)).to.have.property("created.lt");
			expect(qs.parse(v)).to.have.property("created.gte");
		});

		it("inTheLast months", () => {
			var v = encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "inTheLast",
					inTheLastValue: "3",
					inTheLastUnit: "months"
				}
			]);

			expect(qs.parse(v)).to.have.property("created.lt");
			expect(qs.parse(v)).to.have.property("created.gte");
		});

		it("isAfterOrOn", () => {
			encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "isAfterOrOn",
					value: moment("2018-04-09"),
					timezone: "utc"
				}
			]).should.equal("created.gte=1523232000");
		});

		it("isAfter", () => {
			encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "isAfter",
					value: moment("2018-04-09")
				}
			]).should.equal("created.gte=1523289600");
		});

		it("isBefore", () => {
			encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "isBefore",
					value: moment("2018-04-09"),
					timezone: "utc"
				}
			]).should.equal("created.lt=1523232000");
		});

		it("isBeforeOrOn", () => {
			encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "isBeforeOrOn",
					value: moment("2018-04-09"),
					timezone: "utc"
				}
			]).should.equal("created.lt=1523318400");
		});

		it("between", () => {
			encodeFilterData([
				{
					key: "created",
					label: "Created",
					itemType: "DateItem",
					selected: true,
					modifier: "between",
					valueFrom: moment("2018-04-09"),
					valueTo: moment("2018-04-10")
				},
				{
					key: "created1",
					label: "Created1",
					itemType: "DateItem",
					selected: true,
					modifier: "between",
					valueFrom: moment("2018-04-09")
				},
				{
					key: "created2",
					label: "Created2",
					itemType: "DateItem",
					selected: true,
					modifier: "between",
					valueTo: moment("2018-04-09")
				},
				{
					key: "created3",
					label: "Created3",
					itemType: "DateItem",
					selected: true,
					modifier: "between"
				},
				{
					key: "confirmed",
					label: "Confirmed",
					itemType: "DateItem",
					selected: true,
					modifier: "between",
					valueFrom: moment("2018-04-09"),
					valueTo: moment("2018-04-10"),
					timezone: "utc"
				}
			]).should.equal(
				"created.gte=1523203200&created.lte=1523289600&created1.gte=1523203200&created2.lte=1523203200&confirmed.gte=1523232000&confirmed.lte=1523318400"
			);
		});

		describe("encodeFilterData NumberItem", () => {
			it("equals", () => {
				encodeFilterData([
					{
						key: "amount",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "equals",
						value: "12"
					},
					{
						key: "amount1",
						label: "Amount1",
						itemType: "NumberItem",
						selected: true,
						modifier: "equals",
						value: null
					}
				]).should.equal("amount=12");
			});

			it("between", () => {
				encodeFilterData([
					{
						key: "amount",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "between",
						valueFrom: 12,
						valueTo: 24
					},
					{
						key: "amount1",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "between",
						valueTo: 24
					},
					{
						key: "amount2",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "between",
						valueFrom: 12
					},
					{
						key: "amount3",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "between"
					}
				]).should.equal("amount.gte=12&amount.lte=24&amount1.lte=24&amount2.gte=12");
			});

			it("greaterThan", () => {
				encodeFilterData([
					{
						key: "amount",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "greaterThan",
						value: 12,
					}
				]).should.equal("amount.gt=12");
			});

			it("lessThan", () => {
				encodeFilterData([
					{
						key: "amount",
						label: "Amount",
						itemType: "NumberItem",
						selected: true,
						modifier: "lessThan",
						value: 12,
					}
				]).should.equal("amount.lt=12");
			});
		});



		describe("encodeFilterData StringItem", () => {
			it("equals", () => {
				encodeFilterData([
					{
						key: "name",
						label: "Name",
						itemType: "StringItem",
						selected: true,
						modifier: "equals",
						value: "felix[]"
					},
					{
						key: "name1",
						label: "Name",
						itemType: "StringItem",
						selected: true,
						modifier: "equals",
						value: null
					}
				]).should.equal("name=felix%5B%5D");
			});

			it("contains", () => {
				encodeFilterData([
					{
						key: "name",
						label: "Name",
						itemType: "StringItem",
						selected: true,
						modifier: "contains",
						value: "felix[]",
					},
					{
						key: "name",
						label: "Name",
						itemType: "StringItem",
						selected: true,
						modifier: "contains"
					}
				]).should.equal("name.ilike=felix%5B%5D");
			});
		});


	});
});
*/