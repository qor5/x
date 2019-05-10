import BranPopper from './BranPopper';

import { createCustomElement, DOMModel, byContent, byJsonAttrVal } from "@adobe/react-webcomponent";

class MyModel extends DOMModel {
}

class PopperModel extends MyModel {
	@byContent('.bran-popper-content > div') content: any;
	@byJsonAttrVal("react") react: any;
	@byContent('.bran-popper-anchorEl *') anchorEl: any;
}

customElements.define("bran-popper", createCustomElement(BranPopper, PopperModel, "container"))
