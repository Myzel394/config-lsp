// Creates a JSON object in the form of:
// {
//     [<Option name>]: documentation
// }
// 
// Searches for <dl> elements with <dt> and <dd> children based
// on the currently selected element in the Elements tab of the

(() => {
    const content = {}
    let currentOption = ""

    const $elements = $0.querySelectorAll(":scope > dt,dd")

    for (const $element of $elements) {
        switch ($element.tagName) {
            case "DT":
                currentOption = $element.textContent.trim()
                break
            case "DD":
                content[currentOption] = $element.textContent.trim()
                break
        }
    }

    console.log(content)
})()
