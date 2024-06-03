import { useEffect, useState } from "react";
import { PopUpFrame } from "../../components/pop-up/PopUpFrame";
import { DropDownSelect } from "../../components/drop-down/DropDown";

export default function ShareSlideOut({popup}) {

    const [sharedWith, setSharedWith] = useState([])
    const fetchSharedWith = () => {
        // TODO Fetch
        let temp = []
        for(let i = 1; i < 6; i++) {
            temp.push({
                name: 'Name Surname',
                email: 'user' + i + '@email.com',
                premission: i < 3 ? 'EDITOR' : 'VIEWER'
            })
        }
        setSharedWith(temp)
    }


    const [pendingEmail, setPendingEmail] = useState([])
    const [pendingAs, setPendingAs] = useState(0)
    const [pendingShare, setPendingShare] = useState([])
    const addPendingShare = (email) => {
        setPendingShare((prev) => [...prev, email])
        setPendingEmail('')
    }
    const removePendingShare = (email) => {
        setPendingShare((prev) => prev.filter(item => item != email))
    }
    const handleEnterPress = (e) => {
        if (e.code == 'Enter') addPendingShare(pendingEmail)
    }
    const sendRequest = () => {
        // TODO Send premissions added
        setSharedWith((prev) => {
            let add = []
            pendingShare.forEach((user) => {
                console.log(user);
                add.push({
                    name: 'Name Surname',
                    email: user,
                    premission: pendingAs == 0 ? 'EDITOR' : 'VIEWER'
                })
            })
            return [...prev, ...add]
        })
        setPendingEmail('')
        setPendingShare([])
    }

    const premissionLevels = [ {label: 'Viewers', value: 0}, {label: 'Editors', value: 1}]


    useEffect(() => {
        fetchSharedWith()
    }, [])


    return(
        <PopUpFrame visible={popup.isVisible}>
            <div style={{
                    position: "relative"
                }}>
                    <div 
                        className="card"
                        style={{
                            position: "absolute",
                            left: "-8px",
                            top: "7px",
                            width: "29.6%",
                            height: "calc(100vh - 24px)",
                            padding: '14px 16px'
                        }}
                    >
                    <div className="flex space-between center">
                        <p className="section-title" style={{margin:'0'}}></p>    
                        <button className="icon-button small-icon-button" onClick={() => popup.hidePopup()}>
                            <span className="material-symbols-outlined icon">close</span>
                        </button>
                    </div>
                    
                    
                    {sharedWith.length != 0 && <div>
                        <p className="card-title v-spacer-xs hi-spacer-xs">Shared With</p>
                        <div className="flex wrap gap-xxs center v-spacer-l">
                            {sharedWith.map(user => {return(
                                <div className="solid-chip shared-with-chip" style={{height:'48px'}}>
                                    <div className="flex center gap-xs main-content main-content h-100">
                                        <span className="material-symbols-outlined icon small-icon">{user.premission == 'EDITOR'? 'edit' : 'visibility'}</span>
                                        <div className="h-spacer-xs">
                                            <p className="card-body">{user.name}</p>
                                            <p className="card-label neutral">{user.premission}</p>
                                        </div>
                                    </div>
                                    <div className="flex center justify-center gap-xs secondary-content h-100" style={{width:'calc(100% - 8px)', translate:'-9px 0'}}>
                                        <span className="material-symbols-outlined icon small-icon">delete</span>
                                        <p className="button-text">DELETE</p>
                                    </div>
                                </div>
                            )})}
                        </div>
                    </div>}


                    <p className="card-title v-spacer-xs hi-spacer-xs">Share</p>
                    <div className="center gap-xs v-spacer-xs" style={{display:'grid', gridTemplateColumns:'2fr 1fr'}}>
                        <div className="input-wrapper regular-border">
                            <span className="material-symbols-outlined icon input-icon">email</span>
                            <input placeholder="Email" value={pendingEmail} onChange={(e) => setPendingEmail(e.target.value)} onKeyUp={handleEnterPress}/>
                        </div>
                        <DropDownSelect icon={'supervisor_account'} placeholder={'Premission'} options={premissionLevels} callback={(val) => setPendingAs(val)} initialValue={0}/>
                    </div>
                    <div className="flex wrap center gap-xxs v-spacer-s">
                        {pendingShare.map(pending => {return(
                            <div className="info-chip showing flex center" style={{paddingRight:'0'}}>
                                <p className="card-body medium h-spacer-xxs">{pending}</p>
                                <button className="icon-button small-icon-button showing"><span className="material-symbols-outlined icon" onClick={() => removePendingShare(pending)}>cancel</span></button>
                            </div>
                        )})}
                    </div>
                    {pendingShare?.length != 0 && <div className="flex justify-end">
                        <button className="solid-accent-button" onClick={sendRequest}>Add as {premissionLevels.find(item => item.value == pendingAs).label}</button>
                    </div>}

                </div>
            </div>
        </PopUpFrame>
    )
}