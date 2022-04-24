import React, { useState } from 'react'

export default function DNATest() {
    const [patientName, setPatientName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault();
    }
    
    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = (event) => {
        // axios or sequelize
    }

    return (
        <section className="container">
            <h1 className="header">DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Patient name:
                    </label>
                    <input type="text" placeholder="Input the patient name" onChange={e=>setPatientName(e.target.value)}/>
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" onChange={fileInputHandler}/>
                </div>
                <div className="form-control">
                    <label>
                        Disease to predict:
                    </label>
                    <select>
                        <option>Penyakit 1</option>
                        <option>Penyakit 2</option>
                        <option>Penyakit 3</option>
                        <option>Penyakit 4</option>
                    </select>
                </div>

                <input type='submit' value='Submit Test' className = 'btn' onclick={fileSubmitHandler}/>
            </form>
        </section>
    )
}