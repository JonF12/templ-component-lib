import { mergeAttributes, Node } from "@tiptap/core";

export default Node.create({
  name: "newNodeView",
  group: "block",
  atom: true,
  content: "inline*",
  parseHTML() {
    return [
      {
        tag: "new-node-view",
      },
    ];
  },

  renderHTML({ HTMLAttributes }) {
    return ["new-node-view", mergeAttributes(HTMLAttributes)];
  },

  addNodeView() {
    return ({ editor, node, getPos }) => {
      const { view } = editor;

      // Markup
      /*
        <div class="node-view">
          <span class="label">Node view</span>

          <div class="content">
            <button>
              This button has been clicked ${node.attrs.count} times.
            </button>
          </div>
        </div>
      */

      const dom = document.createElement("div");
      dom.classList.add("new-node-view");

      const content = document.createElement("div");
      content.classList.add("content");

      const xhr = new XMLHttpRequest();
      xhr.open("GET", "http://localhost:3000/getheading", false); // false makes it synchronous
      xhr.send();

      if (xhr.status === 200) {
        content.innerHTML = xhr.responseText;
      }

      dom.append(content);

      return {
        dom,
      };
    };
  },
});
