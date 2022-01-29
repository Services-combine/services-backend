import React from 'react';
import classes from './Select.module.css'

const Select = ({options, defaultValue, value, onChange}) => {
    return (
        <select 
            className={classes.Select}
            value={value}
            onChange={event => onChange(event.target.value)}
        >
            <option value="">{defaultValue}</option>
            {options.map(option => 
                <option key={option.value} value={option.value}>
                    {option.name}
                </option>
            )}
        </select>
    );
};

export default Select;