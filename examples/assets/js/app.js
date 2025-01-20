import React from "react";
import { createRoot } from "react-dom/client";
import TipTapEditor from "./tiptapeditor.jsx";

function mountReactComponents() {
  const reactContainers = document.querySelectorAll("[data-react-component]");
  reactContainers.forEach((container) => {
    const componentName = container.dataset.reactComponent;
    const root = createRoot(container);

    switch (componentName) {
      case "tiptapeditor":
        root.render(<TipTapEditor />);
        break;
    }
  });
}
document.addEventListener("DOMContentLoaded", mountReactComponents);
