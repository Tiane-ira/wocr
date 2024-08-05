// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';

export function AddConfig(arg1:model.SkConfig):Promise<boolean>;

export function CheckSk(arg1:model.SkConfig):Promise<string>;

export function DoOcr(arg1:model.OcrParam):Promise<string>;

export function GetConfig():Promise<Array<model.SkConfig>>;

export function GetConfigCount():Promise<number>;

export function GetFields():Promise<Array<model.ExportField>>;

export function OpenFile(arg1:string):Promise<void>;

export function RemoveConfig(arg1:string):Promise<boolean>;

export function SelectDir():Promise<string>;

export function UpdateFields(arg1:Array<number>):Promise<void>;
