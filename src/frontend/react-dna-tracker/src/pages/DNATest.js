export default function DNATest() {
    function handleSubmit(e) {
        e.preventDefault();
    }
    
    return (
        <section className="container">
            <h1 className="header">DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Patient name:
                    </label>
                    <input type="text" placeholder="Input the patient name" />
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file"/>
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

                <input type='submit' value='Submit Test' className = 'btn'/>
            </form>
        </section>
    )
}