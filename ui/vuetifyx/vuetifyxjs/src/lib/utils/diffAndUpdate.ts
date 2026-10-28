/**
 * In-place DOM patch used by ScrollIframe to update the page-builder preview
 * without destroying and rebuilding the whole body (which would reset the iframe
 * scroll position — see KGM-3886).
 *
 * It recursively reconciles `oldNode` to match `newNode`, mutating `oldNode` in
 * place: text and attribute changes are applied to the existing nodes, so any node
 * that did not change keeps its identity (and, in a laid-out document, its size and
 * scroll offset).
 *
 * IMPORTANT: this is an index-based, key-less diff. It is only correct when the two
 * trees have the same child count and order at every level (i.e. no inserts, moves
 * or deletes). If the structure changes, every node after the change point is
 * mis-aligned and patched into the wrong target. Callers must therefore only use it
 * for in-place updates where the structure is guaranteed unchanged (e.g. a page
 * builder field edit, which never adds/moves/removes containers). For structural
 * changes, replace the content wholesale instead.
 */
export const diffAndUpdate = (oldNode: Node, newNode: Node, deep: number = 0) => {
  if (deep > 0) {
    if (oldNode.nodeType !== newNode.nodeType || oldNode.nodeName !== newNode.nodeName) {
      const parent = oldNode.parentNode
      if (parent) {
        parent.replaceChild(newNode.cloneNode(true), oldNode)
      }
      return
    }

    if (oldNode.nodeType === Node.TEXT_NODE) {
      if (oldNode.nodeValue !== newNode.nodeValue) {
        oldNode.nodeValue = newNode.nodeValue
      }
      return
    }
    const oldElement = oldNode as Element
    const newElement = newNode as Element
    const oldAttrs = oldElement.attributes
    const newAttrs = newElement.attributes
    Array.from(oldAttrs).forEach((attr) => {
      if (!newElement.hasAttribute(attr.name)) {
        oldElement.removeAttribute(attr.name)
      }
    })
    Array.from(newAttrs).forEach((attr) => {
      if (oldElement.getAttribute(attr.name) !== attr.value) {
        oldElement.setAttribute(attr.name, attr.value)
      }
    })
  }

  const oldChildren = Array.from(oldNode.childNodes)
  const newChildren = Array.from(newNode.childNodes)
  const maxLength = Math.max(oldChildren.length, newChildren.length)
  for (let i = 0; i < maxLength; i++) {
    if (!oldChildren[i] && newChildren[i]) {
      oldNode.appendChild(newChildren[i].cloneNode(true))
    } else if (oldChildren[i] && !newChildren[i]) {
      oldChildren[i].remove()
    } else if (oldChildren[i] && newChildren[i]) {
      diffAndUpdate(oldChildren[i], newChildren[i], deep + 1)
    }
  }
}
