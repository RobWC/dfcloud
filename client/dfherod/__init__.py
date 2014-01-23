#!/usr/bin/python
import hashlib
import os
import gevent

class FileSet(object):
  def __init__(self,path):
    self.path = path
    self.files = []
  def fileScan(self):
    #find all the files in the specified path
    #add
    totalItems = 0
    totalSize = 0
    itemsInPath = os.listdir(self.path)
    for x in itemsInPath:
      pathToItem = self.path + '/' + x
      if os.path.isdir(pathToItem):
        #recurse
        for (path, dirs, files) in os.walk(pathToItem):
          for filename in files:
            pathToFile = path + '/' + filename
            testfile = File(pathToFile)
            self.files.append(testfile)
            testfile.fileChecksum()
            totalSize = totalSize + testfile.checksum['size']
            totalItems = totalItems + 1;
            
      elif os.path.isfile(pathToItem):
        testfile = File(pathToItem)
        self.files.append(testfile)
        testfile.fileChecksum()
        totalSize = totalSize + testfile.checksum['size']
        totalItems = totalItems + 1;


    print 'Total File Count: {0}\nTotal Size of Files: {1}\n'.format(totalItems,totalSize)


class File(object):
  def __init__(self,path):
    self.path = os.path.abspath(path)
    self.name = 'foo'
    self.fileSize = os.path.getsize(self.path)
    self.checksum = {"value":None,"type":"sha256","size":self.fileSize}
  def readFile(self):
    return open(self.path).read()
  def fileChecksum(self):
    self.checksum['value'] = hashlib.sha256(self.readFile()).hexdigest()

class DFHerod(object):
  def __init__(self):
    pass
  def fileChecksum(self,file):
    print hashlib.sha256(file.readFile()).hexdigest()

fileTestSet = FileSet('/home/rcameron/code/dfcloud/client/dfherod/save')
fileTestSet.fileScan()