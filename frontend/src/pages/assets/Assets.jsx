import React, {useState, useEffect } from "react";
import { getAssets } from "../../api/assets";
import List from "../../components/lists/List";
import SearchBar from "../../components/common/SearchBar";
import { toast } from "react-toastify";

export default function Assets(){
    const [assets, setAssets] = useState();

    const loadAssets = () => {
        const loader = toast.loading("Loading assets...", {
            autoClose: 5000,
            closeButton: true
        })  
        Promise.resolve(getAssets([])).then(data => {

            setAssets(data.list ? data.list : []);
            toast.update(loader, { render: "Fetched assets!", type: "success", isLoading: false, autoClose: 5000, closeButton: true});
        }).catch(error=> {

            setAssets([]);
            toast.update(loader, { render: "Failed to load assets 🤯", type: "error", isLoading: false , autoClose: 5000, closeButton: true});
        });
    }

    useEffect(() => {
        loadAssets();
      }, [])

    return(<div class="grid">
        <div class="text-xl">
        Assets...
        </div>

              
        <SearchBar
        onSubmit={loadAssets}
         />        
     
        <div class="mt-10">
            <List
                columns={[
                    {
                        "header": "AgentId",
                        "colKey": "AgentId"
                    },
                    {
                        "header": "HostName",
                        "colKey": "HostName"
                    },
                    {
                        "header": "Platform",
                        "colKey": "Platform"
                    },
                    {
                        "header": "OperatingSystem",
                        "colKey": "OperatingSystem"
                    },
                    {
                        "header": "Architecture",
                        "colKey": "Architecture"
                    },
                    {
                        "header": "SyncTime",
                        "colKey": "SyncTime"
                    }
                ]}
                data={assets ? assets : []}
            />
        </div>
    </div>)
}