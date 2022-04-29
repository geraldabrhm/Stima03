import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function DNATestResult() {
    const [diseaseName, setDiseaseName] = useState("")
    const [dateTest, setDateTest] = useState()
    function handleSubmit(e) {
        e.preventDefault();
    }

    const fileSubmitHandler = async () => {
        var date_val = document.getElementById("dna_result_date").value
        var disease_name = document.getElementById("disease-name").value
        var text
        var lastEnd = Endpoints.prediction.concat("/" + date_val)
        if(diseaseName == "") {
            const lastRes = await fetch(lastEnd, {
                method: "GET"
            }).then(res => {
                return res.json()
            }).then(data => {
                console.log(data.data)
                for (let i = 0; i < data.data.length; i++) {
                    text += data.data[i].created_at + " - " + data.data[i].patient_name +  "<br>";
                }
            }).catch(err => {
                console.log('ErrorMessage: (Get) ', err)
            })
            
        } else {
            const lastRes = await fetch(lastEnd.concat("/" + disease_name), {
                method: "GET"
            }).then(res => {
                return res.json()
            }).then(data => {
                console.log(data.data)
                for (let i = 0; i < data.data.length; i++) {
                    text += data.data[i].created_at + " - " + data.data[i].patient_name +  "<br>";
                }
            }).catch(err => {
                console.log('ErrorMessage: (Get) ', err)
            })
        }
        document.getElementById("result-test").innerHTML = text
    }

    // const prevResult = () => {
        
    // }

    // const nextResult = () => {

    // }



    return (
        <section className="container2">
            <h1>DNA Test Result</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Date test:
                    </label>
                    <input type="date" id="dna_result_date" onChange={e=>setDateTest(e.target.value)}/>
                </div>
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <input type="text" id="disease-name" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/>
                </div>
                <input type='submit' value='Find Result' className = 'btn' onClick={fileSubmitHandler}/>
            </form>
            <p id="result-test">Result</p>
            {/* <input type='submit' value='Previous' className = 'btn-prev' onClick={nextResult}/>
            <input type='submit' value='Next' className = 'btn-next' onClick={nextResult}/> */}
            
        </section>
    )
}
// TODO get list of DiseaseName from server
// TODO get list of DNATestResult from server