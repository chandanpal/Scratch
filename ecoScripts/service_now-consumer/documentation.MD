## ![alt text](https://raw.githubusercontent.com/ciscoecosystem/Scratch/max-test/ecoScripts/service_now/icon.png "Logo") ServiceNow

### Overview
The ServiceNow integration ...

### Behavior

#### TODO add more

### Configuration

This section desribes how to configure a **deployment** of the ServiceNow image. There are three sections of configuration parameters.

#### ACI Options

- **Tenant Name** is the tenant under which the integration will push segmentation configuration to.
- **Application Profile** is the application profile for the above tenant under which EPGs will be configured.

#### *Renaming attributes

Each parameter in the configuration screen that can be enabled/disabled can also be renamed. For example, the user might want to push `Hostname` to Tetration but search for it using a different parameter name (like `Infoblox Hostname` or `Infoblox_Hostname`) in the Tetration user interface.

It is recommended to maintain consistent names (labels) for annotations in Tetration as Tetration considers `Infoblox Hostname` and `Infoblox_Hostname` to be completely different and unrelated fields.
