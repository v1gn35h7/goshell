import React,  { useEffect, useState } from "react";
import List from "../../components/lists/List";
import SearchBar from "../../components/common/SearchBar";
import { searchResults } from "../../api/scriptApi";
import { toast } from "react-toastify";

export default function Search(){
    const [results, setResults] = useState();


    const loadResults = (query) => {
        const loader = toast.loading("Loading results...", {
            autoClose: 5000,
            closeButton: true
        })  
        Promise.resolve(searchResults({query: query})).then(data => {
            console.log(data);

            setResults(data.list ? data.list : []);
            toast.update(loader, { render: "Fetched results!", type: "success", isLoading: false, autoClose: 5000, closeButton: true});
        }).catch(error=> {
            console.log(error);

            setResults([]);
            toast.update(loader, { render: "Failed to load results 🤯", type: "error", isLoading: false , autoClose: 5000, closeButton: true});
        });
    }

    useEffect(() => {
        loadResults();
      }, [])

    return (
        <div class="grid">
            <div class="text-xl">
            Search...
            </div>

            <SearchBar 
            onSubmit={loadResults}
            />

            <div class="mt-10">
                <List
                    columns={[
                        {
                            "header": "AgentId",
                            "colKey": "agentId"
                        },
                        {
                            "header": "hostname",
                            "colKey": "hostName"
                        },                     
                        {
                            "header": "Output",
                            "colKey": "output"
                        }
                    ]}
                    data={results ? results : []}
                />
            </div>
        </div>
    )
}