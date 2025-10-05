import React,{useEffect, useState} from "react";

export default function List({columns, data}) {

   const dataList = [];

   data.forEach(val => {
        let colData = {};
        columns.forEach(col => {
            colData[col.colKey] = val[col.colKey]
        });
        dataList.push(colData);
    });

    console.log(dataList);

    
    return(
        <div class="grid">
            <table class="border-collapse border border-slate-400 ...">
            <thead>
                <tr>
                    {
                        columns.map(col => <th class="border border-slate-300 ..." align= "left">{col.header}</th>)
                    }              
                </tr>
            </thead>
            <tbody>               
                {
                    dataList.map(val => {
                        return (<tr>{
                            Object.keys(val).map(k => <td class="border border-slate-300 ...">{val[k]}</td>)
                        }</tr>);
                    })
                }    
            </tbody>
            </table>
        </div>
    )
}