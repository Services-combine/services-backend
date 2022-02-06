import React from 'react'
import AccountItem from './AccountItem';

const AccountList = ({accounts, remove}) => {
    return (
        <div className='accounts'>
            {accounts.map((account, index) => 
                <AccountItem remove={remove} account={account} index={index} key={index} />
            )}
        </div>
    );
};

export default AccountList;
