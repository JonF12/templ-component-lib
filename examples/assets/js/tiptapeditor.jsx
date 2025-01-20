import React from "react";
import { useEditor, EditorContent, ReactNodeViewRenderer } from "@tiptap/react";
import { Document } from "@tiptap/extension-document";
import StarterKit from "@tiptap/starter-kit";

// TipTap Extensions
import LinkExtension from "@tiptap/extension-link";
import Typography from "@tiptap/extension-typography";
import TextAlign from "@tiptap/extension-text-align";
import UnderlineExtension from "@tiptap/extension-underline";
import ColorExtension from "@tiptap/extension-color";
import HighlightExtension from "@tiptap/extension-highlight";
import PlaceholderExtension from "@tiptap/extension-placeholder";
import TableExtension from "@tiptap/extension-table";
import TableRow from "@tiptap/extension-table-row";
import TableCell from "@tiptap/extension-table-cell";
import TableHeader from "@tiptap/extension-table-header";
import ImageExtension from "@tiptap/extension-image";
import TaskList from "@tiptap/extension-task-list";
import TaskItem from "@tiptap/extension-task-item";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import { Node } from "@tiptap/core";
import NewNodeView from "./extension";
// Lowlight for code syntax highlighting
import { common, createLowlight } from "lowlight";
// Lucide Icons
import {
  Type,
  Bold,
  Italic,
  Underline as UnderlineIcon,
  Strikethrough,
  List,
  ListOrdered,
  CheckSquare,
  AlignLeft,
  AlignCenter,
  AlignRight,
  Code2,
  Quote,
  Link as LinkIcon,
  Image as ImageIcon,
  Heading1,
  Heading2,
  Heading3,
  LayoutGrid,
} from "lucide-react";

const lowlight = createLowlight(common);

const HeadingComponentNode = Node.create({
  name: "headingComponent",

  group: "block",

  // Allow HTML content inside
  content: "inline*",

  parseHTML() {
    return [
      {
        tag: "header.mb-8",
        getContent: (element) => element.innerHTML,
        // This tells TipTap how to parse the incoming HTML
        parseHTML: true,
      },
    ];
  },

  renderHTML({ node }) {
    // Directly use the HTML content
    return [
      "div",
      {
        "data-type": "heading-component",
        "data-component": true,
        // Use dangerouslySetInnerHTML equivalent in TipTap
        innerHTML: node.content,
      },
    ];
  },
});

const MenuButton = ({ editor, onClick, isActive = false, title, children }) => {
  return (
    <button
      onClick={onClick}
      className={`p-2 rounded hover:bg-gray-100 ${isActive ? "bg-gray-200" : ""}`}
      title={title}
    >
      {React.cloneElement(children, { className: "w-5 h-5" })}
    </button>
  );
};

const MenuBarGroup = ({ children, isLast = false }) => {
  return (
    <div className={`flex gap-1 ${isLast ? "" : "border-r pr-2"}`}>
      {children}
    </div>
  );
};

const MenuBar = ({ editor }) => {
  if (!editor) {
    return null;
  }

  return (
    <div className="border-b p-2 flex flex-wrap gap-2">
      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() =>
            editor.chain().focus().toggleHeading({ level: 1 }).run()
          }
          isActive={editor.isActive("heading", { level: 1 })}
          title="Heading 1"
        >
          <Heading1 />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() =>
            editor.chain().focus().toggleHeading({ level: 2 }).run()
          }
          isActive={editor.isActive("heading", { level: 2 })}
          title="Heading 2"
        >
          <Heading2 />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() =>
            editor.chain().focus().toggleHeading({ level: 3 }).run()
          }
          isActive={editor.isActive("heading", { level: 3 })}
          title="Heading 3"
        >
          <Heading3 />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleBold().run()}
          isActive={editor.isActive("bold")}
          title="Bold"
        >
          <Bold />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleItalic().run()}
          isActive={editor.isActive("italic")}
          title="Italic"
        >
          <Italic />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleUnderline().run()}
          isActive={editor.isActive("underline")}
          title="Underline"
        >
          <UnderlineIcon />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleStrike().run()}
          isActive={editor.isActive("strike")}
          title="Strikethrough"
        >
          <Strikethrough />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleBulletList().run()}
          isActive={editor.isActive("bulletList")}
          title="Bullet List"
        >
          <List />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleOrderedList().run()}
          isActive={editor.isActive("orderedList")}
          title="Ordered List"
        >
          <ListOrdered />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleTaskList().run()}
          isActive={editor.isActive("taskList")}
          title="Task List"
        >
          <CheckSquare />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().setTextAlign("left").run()}
          isActive={editor.isActive({ textAlign: "left" })}
          title="Align Left"
        >
          <AlignLeft />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().setTextAlign("center").run()}
          isActive={editor.isActive({ textAlign: "center" })}
          title="Align Center"
        >
          <AlignCenter />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().setTextAlign("right").run()}
          isActive={editor.isActive({ textAlign: "right" })}
          title="Align Right"
        >
          <AlignRight />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleCodeBlock().run()}
          isActive={editor.isActive("codeBlock")}
          title="Code Block"
        >
          <Code2 />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => editor.chain().focus().toggleBlockquote().run()}
          isActive={editor.isActive("blockquote")}
          title="Quote"
        >
          <Quote />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup>
        <MenuButton
          editor={editor}
          onClick={() => {
            const url = window.prompt("Enter URL:");
            if (url) {
              editor.chain().focus().setLink({ href: url }).run();
            }
          }}
          isActive={editor.isActive("link")}
          title="Add Link"
        >
          <LinkIcon />
        </MenuButton>
        <MenuButton
          editor={editor}
          onClick={() => {
            const url = window.prompt("Enter image URL:");
            if (url) {
              editor.chain().focus().setImage({ src: url }).run();
            }
          }}
          title="Add Image"
        >
          <ImageIcon />
        </MenuButton>
      </MenuBarGroup>

      <MenuBarGroup isLast>
        <MenuButton
          editor={editor}
          onClick={() => {
            editor.commands.insertContent("<new-node-view></new-node-view>");
          }}
          title="blah"
        >
          <LayoutGrid />
        </MenuButton>
      </MenuBarGroup>
    </div>
  );
};

const TipTapEditor = ({ initialContent = "", onUpdate }) => {
  const editor = useEditor({
    content: `
<p> test </p>
<new-node-view> 

</new-node-view>`,
    extensions: [
      StarterKit.configure({
        heading: {
          levels: [1, 2, 3],
        },
        codeBlock: false,
      }),
      // Enable HTML parsing
      Document.extend({
        content: "block+",
      }),
      LinkExtension.configure({
        openOnClick: false,
        HTMLAttributes: {
          class: "text-blue-500 underline",
        },
      }),
      Typography,
      TextAlign.configure({
        types: ["heading", "paragraph"],
      }),
      UnderlineExtension,
      ColorExtension,
      HighlightExtension.configure({
        multicolor: true,
      }),
      PlaceholderExtension.configure({
        placeholder: "Write something...",
      }),
      TableExtension.configure({
        resizable: true,
      }),
      TableRow,
      TableCell,
      TableHeader,
      ImageExtension.configure({
        HTMLAttributes: {
          class: "max-w-full h-auto",
        },
      }),
      TaskList,
      TaskItem.configure({
        nested: true,
      }),
      HeadingComponentNode,
      NewNodeView,
      CodeBlockLowlight.configure({
        lowlight: lowlight,
        defaultLanguage: "javascript",
      }),
    ],
    content: initialContent,
    onUpdate: ({ editor }) => {
      const html = editor.getHTML();
      if (onUpdate) {
        onUpdate(html);
      }
    },
  });

  return (
    <div className="w-full border rounded-lg overflow-hidden">
      <MenuBar editor={editor} />
      <EditorContent
        editor={editor}
        className="prose max-w-none p-4 min-h-lvh [&_.ProseMirror-focused]:outline-none"
      />
    </div>
  );
};

export default TipTapEditor;
