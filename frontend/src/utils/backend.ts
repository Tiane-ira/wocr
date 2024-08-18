import { SelectDir, DoOcr, AddConfig, RemoveConfig, GetConfig, GetConfigCount, CheckSk, GetFields, UpdateFields,ChangeDefault } from "../../wailsjs/go/service/App";
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime';

export const handleSelectDir = async () => {
    return await SelectDir();
};

export const startOcr = async (param) => {
    return await DoOcr(param);
};

export const getConfigs = async () => {
    return await GetConfig();
};


export const saveConfig = async (param) => {
    return await AddConfig(param);
};

export const removeConfig = async (id) => {
    return await RemoveConfig(id);
};

export const getConfigCount = async () => {
    return await GetConfigCount();
};

export const checkConfig = async (param) => {
    return await CheckSk(param);
};

export const getFields = async () => {
    return await GetFields();
}

export const updateFields = async (idList) => {
    return await UpdateFields(idList);
}

export const setDefaultConfig = async (id) => {
    return await ChangeDefault(id);
}

export const eventLogOn = (fn) => {
    EventsOn("ocr_log", fn)
};

export const eventLogOff = () => {
    EventsOff("ocr_log")
};

