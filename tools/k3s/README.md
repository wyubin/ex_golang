# intro
紀錄安裝k3s 所需要的相關套件

# condition
確認虛擬化
```shell
sudo grep --color -E "vmx|svm" /proc/cpuinfo
# check linux lib
sudo lsmod | grep kvm
```

# tool install

## qemu
參考 [link](https://ivonblog.com/posts/archlinux-qemu-virt-manager/)
```shell
sudo pacman -S 
# install cloud-init
sudo pacman -S cloud-init cloud-utils
```

# create vm
## prepare vm
參考 [Create Debian 12 (Bookworm) KVM Guest From Cloud Image](https://blog.programster.org/create-debian-12-kvm-guest-from-cloud-image)
- download cloud image
```shell
wget https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-generic-amd64.qcow2

# resise a bigger image
cp debian-12-generic-amd64.qcow2 vm-debian-1.qcow2
qemu-img info vm-debian-1.qcow2
qemu-img resize vm-debian-1.qcow2 4G
```

## create config image
因為 could config 是以一個 storage 的方式放進 vm, 所以要先建立 config image
```shell
cloud-localds ./vm1-config.img ./cloud-config.yaml
```

## create guest
```shell
qemu-system-x86_64 \
  -m 512M -hda vm-debian-1.qcow2 \
  -hdb vm1-config.img -net user -net nic

sudo virt-install \
  --name $VM_NAME \
  --memory 1024 \
  --disk vm-debian-1.qcow2,device=disk,bus=virtio \
  --disk vm1-config.img,device=cdrom \
  --os-type linux \
  --os-variant debian10 \
  --virt-type kvm \
  --graphics none \
  --network network=c,model=virtio \
  --import
```