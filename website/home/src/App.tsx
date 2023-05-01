import React from "react";
import {createRoot} from "react-dom/client";

import "./index.scss";

// @ts-ignore
import Display from "display/Display";
// @ts-ignore
import Detail from "detail/Detail";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";

const App = () => (
    <>
        <div className="container mx-auto mt-8">
            <h1 className="text-3xl font-bold mb-8">Home page</h1>
        </div>

        <Router>
            <div>
                <Routes>
                    <Route path="/" element={<Display/>}/>
                    <Route path="/:isbn" element={<Detail/>}/>
                </Routes>
            </div>
        </Router>

    </>
);

const rootElement = document.getElementById("app")!;
const root = createRoot(rootElement);
root.render(<App/>);
