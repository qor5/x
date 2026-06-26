import { describe, it, expect } from 'vitest'
import { diffAndUpdate } from './diffAndUpdate'

// Build a detached <body> with the given inner HTML (no inter-tag whitespace so the
// index-based diff aligns deterministically).
const makeBody = (html: string): HTMLElement => {
  const body = document.createElement('body')
  body.innerHTML = html
  return body
}

const containers = (body: HTMLElement) =>
  Array.from(body.querySelectorAll('[data-container-id]')) as HTMLElement[]

describe('diffAndUpdate (ScrollIframe in-place patch, KGM-3886)', () => {
  it('preserves the identity of unchanged container nodes when one field changes', () => {
    // This is the property that fixes KGM-3886: a field edit re-renders the whole
    // body, but if unchanged container DOM nodes keep their identity, their laid-out
    // size is preserved and the iframe scroll position cannot move — no jump.
    const oldBody = makeBody(
      '<div data-container-id="a"><p>Alpha</p></div>' +
        '<div data-container-id="b"><p>Beta</p></div>' +
        '<div data-container-id="c"><p>Gamma</p></div>'
    )
    const newBody = makeBody(
      '<div data-container-id="a"><p>Alpha</p></div>' +
        '<div data-container-id="b"><p>Beta EDITED</p></div>' +
        '<div data-container-id="c"><p>Gamma</p></div>'
    )

    const [aBefore, bBefore, cBefore] = containers(oldBody)

    diffAndUpdate(oldBody, newBody)

    const [aAfter, bAfter, cAfter] = containers(oldBody)
    // Unchanged containers: same node objects (not rebuilt).
    expect(aAfter).toBe(aBefore)
    expect(cAfter).toBe(cBefore)
    // Edited container: same node, only its inner text patched.
    expect(bAfter).toBe(bBefore)
    expect(bAfter.textContent).toBe('Beta EDITED')
    expect(aAfter.textContent).toBe('Alpha')
  })

  it('patches a changed attribute in place without replacing the node', () => {
    // e.g. a Video container whose embed src changes: only the existing node's src is
    // updated (so only that one embed reloads), stale attributes are dropped, and the
    // node identity — and therefore every other node's layout — is untouched.
    const oldBody = makeBody(
      '<div data-container-id="v"><iframe src="https://youtube.com/embed/OLD" data-stale="1"></iframe></div>'
    )
    const newBody = makeBody(
      '<div data-container-id="v"><iframe src="https://youtube.com/embed/NEW" title="t"></iframe></div>'
    )

    const iframeBefore = oldBody.querySelector('iframe')!

    diffAndUpdate(oldBody, newBody)

    const iframeAfter = oldBody.querySelector('iframe')!
    expect(iframeAfter).toBe(iframeBefore)
    expect(iframeAfter.getAttribute('src')).toBe('https://youtube.com/embed/NEW')
    expect(iframeAfter.getAttribute('title')).toBe('t')
    expect(iframeAfter.hasAttribute('data-stale')).toBe(false)
  })

  it('replaces a node whose tag changed', () => {
    const oldBody = makeBody('<div data-container-id="a"><span>x</span></div>')
    const newBody = makeBody('<div data-container-id="a"><p>x</p></div>')

    diffAndUpdate(oldBody, newBody)

    expect(oldBody.querySelector('[data-container-id="a"] p')).not.toBeNull()
    expect(oldBody.querySelector('[data-container-id="a"] span')).toBeNull()
  })

  it('documents the boundary: a structural change scrambles node identity', () => {
    // Inserting a container at the front shifts every sibling. The index-based,
    // key-less diff has no way to know, so it rewrites each existing node to look
    // like the next one and clones a new tail node. The final DOM is content-correct,
    // but every node after the change point was needlessly rewritten (identity lost
    // → real embeds would reload, scroll would not be preserved). This is exactly why
    // ScrollIframe only enables diffAndUpdate for isUpdate === true (field edits,
    // structure unchanged) and keeps innerHTML for add/move/delete.
    const oldBody = makeBody(
      '<div data-container-id="a"></div><div data-container-id="b"></div>'
    )
    const newBody = makeBody(
      '<div data-container-id="new"></div>' +
        '<div data-container-id="a"></div>' +
        '<div data-container-id="b"></div>'
    )

    const aBefore = oldBody.querySelector('[data-container-id="a"]')

    diffAndUpdate(oldBody, newBody)

    // Content is correct...
    expect(containers(oldBody).map((c) => c.getAttribute('data-container-id'))).toEqual([
      'new',
      'a',
      'b'
    ])
    // ...but the node that was "a" is now relabeled "new" — identity scrambled.
    expect(aBefore!.getAttribute('data-container-id')).toBe('new')
  })
})
