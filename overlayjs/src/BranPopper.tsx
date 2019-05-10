import React, { FC, useState } from 'react';
import Popper from '@material-ui/core/Popper';

const BranPopper: FC = (props: any) => {
	const [open, setOpen] = useState(false)
	const branProps = props.react
	const content = props.content
	var anchorEl = props.anchorEl
	console.log("props.content", props.content)
	console.log("anchorEl", anchorEl)
	// console.log("anchorEl.props.item.node", anchorEl.props.item.node)
	anchorEl.onClick = (e: Event) => {
		console.log("hi")
		setOpen(!open)
	}

	return (
		<div className="abc">
			{anchorEl}

			<Popper open={open} {...branProps}>
				{content}
			</Popper >
		</div>
	);
}

export default BranPopper;
