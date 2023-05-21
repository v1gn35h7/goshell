import React,  { useEffect, useState } from "react";
import List from "../../components/lists/List";
import SearchBar from "../../components/common/SearchBar";
import {  Link } from "react-router-dom";
import { getScripts } from "../../api/scriptApi";
import { toast } from "react-toastify";
import { PlusIcon } from "@heroicons/react/24/solid";

export default function Scripts() {
    const [scripts, setScripts] = useState();

    const loadScripts = () => {
        const loader = toast.loading("Loading scripts...", {
            autoClose: 5000,
            closeButton: true
        })  
        Promise.resolve(getScripts([])).then(data => {
            console.log(data);

            setScripts(data.list ? data.list : []);
            toast.update(loader, { render: "Fetched scripts!", type: "success", isLoading: false, autoClose: 5000, closeButton: true});
        }).catch(error=> {
            console.log(error);

            setScripts([]);
            toast.update(loader, { render: "Failed to load scripts 🤯", type: "error", isLoading: false , autoClose: 5000, closeButton: true});
        });
    }

    useEffect(() => {
        loadScripts();
      }, [])

    return (
        <div class="grid">
        <div class="text-xl">
            Scripts...
        </div>
        
        <SearchBar
        onSubmit={loadScripts}
         />
        
        <div class="mt-6">
            <Link class="bg-transparent text-blue-700 font-semibold  py-2 px-4 border border-blue-500  rounded  inline-flex items-center" to={'/scripts/new'}>
                <PlusIcon className="h-6 w-6 text-blue-500" /> Add Script
            </Link>
        </div>

        <div class="mt-10">
            <List
                columns={[
                    {
                        "header": "Id",
                        "colKey": "Id"
                    },
                    {
                        "header": "Platform",
                        "colKey": "Platform"
                    },
                    {
                        "header": "Type",
                        "colKey": "Type"
                    },
                    {
                        "header": "Actions",
                        "colKey": "id"
                    }
                ]}
                data={scripts ? scripts : []}
            />
        </div>
    </div>
    )
}