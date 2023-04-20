package login

import (
	"fmt"
	"net/http"

	. "github.com/theplant/htmlgo"
)

var DefaultViewCommon = &ViewCommon{
	WrapperClass: "flex pt-16 flex-col max-w-md mx-auto",
	TitleClass:   "leading-tight text-3xl mt-0 mb-6",
	LabelClass:   "block mb-2 text-sm text-gray-600 dark:text-gray-200",
	InputClass:   "block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40",
	ButtonClass:  "w-full px-6 py-3 tracking-wide text-white transition-colors duration-200 transform bg-blue-500 rounded-md hover:bg-blue-400 focus:outline-none focus:bg-blue-400 focus:ring focus:ring-blue-300 focus:ring-opacity-50",
}

type ViewCommon struct {
	WrapperClass string
	TitleClass   string
	LabelClass   string
	InputClass   string
	ButtonClass  string
}

func (vc *ViewCommon) Notice(vh *ViewHelper, msgr *Messages, w http.ResponseWriter, r *http.Request) HTMLComponent {
	var nn HTMLComponent
	if n := vh.GetNoticeFlash(w, r); n != nil && n.Message != "" {
		switch n.Level {
		case NoticeLevel_Info:
			nn = vc.InfoNotice(n.Message)
		case NoticeLevel_Warn:
			nn = vc.WarnNotice(n.Message)
		case NoticeLevel_Error:
			nn = vc.ErrNotice(n.Message)
		}
	}
	return Components(
		vc.ErrNotice(vh.GetFailFlashMessage(msgr, w, r)),
		vc.WarnNotice(vh.GetWarnFlashMessage(msgr, w, r)),
		vc.InfoNotice(vh.GetInfoFlashMessage(msgr, w, r)),
		nn,
	)
}

func (vc *ViewCommon) ErrNotice(msg string) HTMLComponent {
	if msg == "" {
		return nil
	}

	return Div().Class("bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative text-center").
		Role("alert").
		Children(
			Span(msg).Class("block sm:inline"),
		)
}

func (vc *ViewCommon) WarnNotice(msg string) HTMLComponent {
	if msg == "" {
		return nil
	}

	return Div().Class("bg-orange-100 border border-orange-400 text-orange-700 px-4 py-3 rounded relative text-center").
		Role("alert").
		Children(
			Span(msg).Class("block sm:inline"),
		)
}

func (vc *ViewCommon) InfoNotice(msg string) HTMLComponent {
	if msg == "" {
		return nil
	}

	return Div().Class("bg-blue-100 border border-blue-400 text-blue-700 px-4 py-3 rounded relative text-center").
		Role("alert").
		Children(
			Span(msg).Class("block sm:inline"),
		)
}

func (vc *ViewCommon) ErrorBody(msg string) HTMLComponent {
	return Div(
		Text(msg),
	)
}

func (vc *ViewCommon) PasswordInputWithRevealFunction(
	name string,
	placeholder string,
	id string,
	val string,
) HTMLComponent {
	return Div(
		Input(name).Placeholder(placeholder).Type("password").Class(vc.InputClass).Class("pr-10").Id(id).
			Value(val),
		Div(
			RawHTML(fmt.Sprintf(`<svg class="h-6 text-gray-700 block" id="icon-%s-showed" fill="none" xmlns="http://www.w3.org/2000/svg" viewbox="0 0 576 512" width="1rem">
  <path fill="currentColor"
    d="M572.52 241.4C518.29 135.59 410.93 64 288 64S57.68 135.64 3.48 241.41a32.35 32.35 0 0 0 0 29.19C57.71 376.41 165.07 448 288 448s230.32-71.64 284.52-177.41a32.35 32.35 0 0 0 0-29.19zM288 400a144 144 0 1 1 144-144 143.93 143.93 0 0 1-144 144zm0-240a95.31 95.31 0 0 0-25.31 3.79 47.85 47.85 0 0 1-66.9 66.9A95.78 95.78 0 1 0 288 160z">
  </path>
</svg>`, id)),
			RawHTML(fmt.Sprintf(`<svg class="h-6 text-gray-700 hidden" id="icon-%s-hidden" fill="none" xmlns="http://www.w3.org/2000/svg" viewbox="0 0 640 512" width="1rem">
  <path fill="currentColor"
    d="M320 400c-75.85 0-137.25-58.71-142.9-133.11L72.2 185.82c-13.79 17.3-26.48 35.59-36.72 55.59a32.35 32.35 0 0 0 0 29.19C89.71 376.41 197.07 448 320 448c26.91 0 52.87-4 77.89-10.46L346 397.39a144.13 144.13 0 0 1-26 2.61zm313.82 58.1l-110.55-85.44a331.25 331.25 0 0 0 81.25-102.07 32.35 32.35 0 0 0 0-29.19C550.29 135.59 442.93 64 320 64a308.15 308.15 0 0 0-147.32 37.7L45.46 3.37A16 16 0 0 0 23 6.18L3.37 31.45A16 16 0 0 0 6.18 53.9l588.36 454.73a16 16 0 0 0 22.46-2.81l19.64-25.27a16 16 0 0 0-2.82-22.45zm-183.72-142l-39.3-30.38A94.75 94.75 0 0 0 416 256a94.76 94.76 0 0 0-121.31-92.21A47.65 47.65 0 0 1 304 192a46.64 46.64 0 0 1-1.54 10l-73.61-56.89A142.31 142.31 0 0 1 320 112a143.92 143.92 0 0 1 144 144c0 21.63-5.29 41.79-13.9 60.11z">
  </path>
</svg>`, id)),
		).Class("absolute right-0 inset-y-0 px-2 flex items-center text-sm cursor-pointer").Id(fmt.Sprintf("btn-reveal-%s", id)),
		Script(fmt.Sprintf(`
(function(){
    var passElem = document.getElementById("%s");
    var revealBtn = document.getElementById("btn-reveal-%s");
    var showedIcon = document.getElementById("icon-%s-showed");
    var hiddenIcon = document.getElementById("icon-%s-hidden");
    revealBtn.onclick = function() {
        if (passElem.type === "password") {
            passElem.type = "text";
            showedIcon.classList.remove("block");
            showedIcon.classList.add("hidden");
            hiddenIcon.classList.remove("hidden");
            hiddenIcon.classList.add("block");
        } else {
            passElem.type = "password";
            hiddenIcon.classList.remove("block");
            hiddenIcon.classList.add("hidden");
            showedIcon.classList.remove("hidden");
            showedIcon.classList.add("block");
        }
    };
})();`, id, id, id, id)),
	).Class("relative")
}

func (vc *ViewCommon) PasswordStrengthMeter(inputID string) HTMLComponent {
	meterID := fmt.Sprintf("%s-strength-meter", inputID)
	return Div(
		Div(
			Div(
				Div().Class("password-strength-meter-section h-2 rounded-xl transition-colors bg-gray-200"),
			).Class("w-1/5 px-1"),
			Div(
				Div().Class("password-strength-meter-section h-2 rounded-xl transition-colors bg-gray-200"),
			).Class("w-1/5 px-1"),
			Div(
				Div().Class("password-strength-meter-section h-2 rounded-xl transition-colors bg-gray-200"),
			).Class("w-1/5 px-1"),
			Div(
				Div().Class("password-strength-meter-section h-2 rounded-xl transition-colors bg-gray-200"),
			).Class("w-1/5 px-1"),
			Div(
				Div().Class("password-strength-meter-section h-2 rounded-xl transition-colors bg-gray-200"),
			).Class("w-1/5 px-1"),
		).Class("flex mt-2 -mx-1 hidden").Id(meterID),
		Script(fmt.Sprintf(`
(function(){
    var passElem = document.getElementById("%s");
    var meterElem = document.getElementById("%s");
    var meterSectionElems = document.getElementsByClassName("password-strength-meter-section");
    function checkStrength(val) {
        if (!val) {
            return 0;
        };
        return zxcvbn(val).score + 1;
    };
    function updateMeter() {
        if (passElem.value) {
            meterElem.classList.remove("hidden");
        } else {
            if (!meterElem.classList.contains("hidden")) {
                meterElem.classList.add("hidden");
            }
        }
        var s = checkStrength(passElem.value);
        for (var i = 0; i < meterSectionElems.length; i++) {
            var elem = meterSectionElems[i];
            if (i >= s) {
                elem.classList.add("bg-gray-200");
                elem.classList.remove("bg-red-400", "bg-yellow-400", "bg-green-500");
            } else if (s <= 2) {
                elem.classList.add("bg-red-400");
                elem.classList.remove("bg-gray-200", "bg-yellow-400", "bg-green-500");
            } else if (s <= 4) {
                elem.classList.add("bg-yellow-400");
                elem.classList.remove("bg-red-400", "bg-gray-200", "bg-green-500");
            } else {
                elem.classList.add("bg-green-500");
                elem.classList.remove("bg-red-400", "bg-yellow-400", "bg-gray-200");
            }
        }
    };
    updateMeter();
    passElem.oninput = function(e) {
        updateMeter();
    };
})();`, inputID, meterID)),
	)
}
