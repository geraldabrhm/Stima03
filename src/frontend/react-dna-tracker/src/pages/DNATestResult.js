export default function DNATestResult() {
    function handleSubmit(e) {
        e.preventDefault();
    }

    const fileSubmitHandler = (event) => {
        // axios or sequelize
    }

    return (
        <section className="container2">
            <h1>DNA Test Result</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Date test:
                    </label>
                    <input type="date" placeholder="Input the disease name" />
                </div>
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <select>
                        <option>Penyakit 1</option>
                        <option>Penyakit 2</option>
                        <option>Penyakit 3</option>
                        <option>Penyakit 4</option>
                    </select>
                </div>
                <input type='submit' value='Find Result' className = 'btn' onClick={fileSubmitHandler}/>
            </form>
            <ul>
                <li>A</li>
                <li>B</li>
                <li>C</li>
                <li>D</li>
                <li>E</li>
            </ul>
            
        </section>
    )
}