import React, {useState} from "react";
import { MagnifyingGlassCircleIcon, PlusIcon } from '@heroicons/react/24/solid';


export default function SearchBar({onSubmit}) {
    const [query, setQuery] = useState("");
    return (
    <div class="grid  mt-10">
        <div class="flex flex-row">
            <div class="basis-3/4">
                <input type="text" class="form-input w-full" placeholder="Search ..." onChange={(e) => setQuery(e.target.value)} />
            </div>
            <div class="basis-1/4">
                <button class="bg-transparent text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 rounded" 
                onClick={() => onSubmit(query) }
                >
                 <MagnifyingGlassCircleIcon className="h-6 w-6 text-blue-500"  />
                </button>
            </div>
        </div>
    </div>
    )
}