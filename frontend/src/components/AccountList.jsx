import React from 'react'
import AccountItem from './AccountItem';

const AccountList = ({accounts}) => {
    return (
        <div className='accounts'>
            {accounts.map((account, index) => 
                <AccountItem account={account} index={index} key={account.id} />
            )}
        </div>
    );
};

export default AccountList;
