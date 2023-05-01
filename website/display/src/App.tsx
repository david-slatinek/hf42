import React from "react";

import "./index.scss";
import Display from "./Display";
import {createRoot} from "react-dom/client";

const App = () => (
    <div>
        <Display/>
    </div>
);

const rootElement = document.getElementById("app")!;
const root = createRoot(rootElement);
root.render(<App/>);
