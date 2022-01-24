import React, { useState } from 'react'

const InputControl = function() {
    const [value, setValue] = useState('text')

    return (
        <div>
            <h1>{value}</h1>
			<input type='text' value={value} onChange={event => setValue(event.target.value)} />
        </div>
    );
};

export default InputControl;
