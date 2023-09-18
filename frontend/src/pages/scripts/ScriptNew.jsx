import React, {useState} from "react";
import { saveScript } from "../../api/scriptApi";
import { toast } from "react-toastify";
import { Link, useNavigate } from "react-router-dom";

export default function ScriptNew(){
    const [title, setTitle] = useState("");
    const [platform, setPlatform] = useState("Windows");
    const [frequency, setFrequency] = useState("* * * * *");
    const [script, setScript] = useState("");
    const navigate = useNavigate();

    const submit = () => {
        toast.promise(
            saveScript({
                title: title,
                platform: platform,
                frequency: frequency,
                script: script
            }),
            {
              pending: 'Saving script!',
              success: 'Saved script 👌',
              error: 'Failed to save 🤯'
            }
        );
        navigate("/scripts");
    };


    return (
        <div>
            <form onSubmit={(e) =>{ 
                e.preventDefault();
                submit();
                 }}>
        <div class="my-6 text-xl">
            New Script
        </div>
        <div class="grid">
            <div class="flex flex-row mt-10">
                <div class="w-full">
                    <input type="text" class="form-input w-full" placeholder="Script Title" required onChange={(e) => {
                        setTitle(e.target.value)
                    }} />
                </div>               
            </div>
            <div class="flex flex-row mt-10">
                <div class="basis-3/4">
                    <select type="text" class="form-input w-full" placeholder="Script Platform" required onChange={(e) => {
                        setPlatform(e.value)
                    }}>
                        <option selected>Windows</option>
                        <option>Linux</option>
                        <option>Darwin</option>
                    </select>
                </div> 
                <div class="basis-3/4">
                    <input type="text" class="form-textarea w-full" placeholder="Frequency (Ex: * * * * *)" required onChange={(e) => {
                        setFrequency(e.target.value)
                    }} />
                </div>               
            </div>
            <div class="flex flex-row mt-10">
                <div class="w-full">
                    <textarea type="text" class="form-textarea w-full" style={{height: "300px"}} placeholder="Enter Script ..." required  onChange={(e) => {
                        setScript(e.target.value)
                    }} />
                </div>               
            </div>
            <div class="flex flex-row mt-10">
                <button type="submit" class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded mr-5">
                    Save
                </button>              
                <Link class="bg-transparent text-blue-700 font-semibold  py-2 px-4 border border-blue-500  rounded  inline-flex items-center" to={'/scripts'}>
                 Cancel
                </Link>
            </div>
        </div>
        </form>
        </div>
    )
}