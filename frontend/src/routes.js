import React from "react";
import Dashboard from "./pages/dashboard/Dashboard";
import Assets from "./pages/assets/Assets";
import Search from "./pages/search/Search";
import Scripts from "./pages/scripts/Scripts";
import ScriptNew from "./pages/scripts/ScriptNew";

var routes = [
    {
        path: "/",
        loader: () => ({ message: "Hello, welcome to GoShell!" }),
        element: <Search />,
      },
      {
        path: "/assets",
        element: <Assets />,
      },
      {
        path: "/search",
        element: <Search />,
      },
      {
        path: "/scripts",
        element: <Scripts />,
        children:[
         
        ]
      },
      {
        path: "/scripts/new",
        element: <ScriptNew />,
      },
      {
        path: "/scripts/{id}",
        element: <ScriptNew />,
      }
];

export default routes;