// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {variables} from '../models';

export function GetWebsites(arg1:{[key: string]: Array<string>}):Promise<Array<string>>;

export function ReadUserSettings():Promise<main.UserSettings>;

export function Search(arg1:string,arg2:Array<string>):Promise<Array<variables.ItemList>>;

export function UpdateUserSettings(arg1:main.UserSettings):Promise<void>;
